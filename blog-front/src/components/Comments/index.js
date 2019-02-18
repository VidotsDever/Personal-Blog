import React, {Component} from 'react';
import axios from 'axios';
import { convquery } from '../../utils/str';
import { getLocalTime } from '../../utils/date';
import Cookie from 'js-cookie';
import qs from 'qs';
import {Icon, Col, Row, Input, Button} from 'antd';
const {TextArea}  = Input




export default class Comments extends Component {

    constructor(props) {
        super(props);
        this.state = {
            name: "",
            avatar: "",
            comment: "",
            comment_behind: "",
            comments: [],
            parent_id: 0,
            replyer_id: "",
            replyer_name: ""
        };
    }

    _resetComment = () => {
        this.setState({
            parent_id: 0,
            replyer_id: "",
            replyer_name: "",
            comment: "",
            comment_behind: ""
        })
    }

    _setComment = (parent_id, replyer_id, replyer_name) => {
        this.setState({
            parent_id: parent_id,
            replyer_id: replyer_id,
            replyer_name: replyer_name
        })
    }

    _isLogged = () => {
        return (this.state.avatar !== "" && this.state.name !== "")
    }

    _handleTextChange = (e) => {
        this.setState({
            comment: e.target.value,
            comment_behind: e.target.value
        });
    }

    _getComments = async () => {
        axios.get(`http://localhost:8080/comment/list?post_id=${this.props.post_id}`).then(res => {
        console.log(res.data.comments)    
        this.setState({
                comments: res.data.comments
            })
        }).catch(err => {
            console.log(err)
        })
    }

    componentDidMount() {
        this._getComments()
    }

    _saveComment = async () => {
        if(this.state.comment === "") {
            alert("评论不能为空")
            return
        }
        const data = {
            content: this.state.comment,
            post_id: this.props.post_id,
            commenter_id: Cookie.get("id") || "",
            replyer_id: this.state.replyer_id,
            parent_id: this.state.parent_id,
            replyer_name: this.state.replyer_name
        };
        this._resetComment()
        axios.post(
            "http://localhost:8080/comment/save",
            qs.stringify(data),
            {
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                  }
            }
        ).then(response => {
            if(response.statusText === "OK") {
                this._getComments()
            } else {
                console.log(response)
            }
        }).catch(err => {
            console.log(err)
        }); 

    }

    _requestUser = async (code) => {
        const data = {
            code: code
        };
        axios.post(
            "http://localhost:8080/auth/user",
            qs.stringify(data),
            {
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                  }
            }
        ).then(response => {
            if(response.statusText === "OK") {
                const id = response.data.user.id
                const avatar = response.data.user.avatar
                const name = response.data.user.name
                Cookie.set("id", id)
                Cookie.set("avatar", avatar)
                Cookie.set("name", name)
                this.setState({
                    avatar: avatar,
                    name: name
                })
            } else {
                console.log(response)
            }
        }).catch(err => {
            console.log(err)
        });
    }

    _requestAuthorization = () => {
        let pop = window.open(`https://github.com/login/oauth/authorize?client_id=b35c9c1d9c1649803683`, null, "width=600,height=400")
        let checkCode = () => {
            try {
                let query = pop.location.search.substr(1);
                let code = convquery(query).code;
                if((typeof code) !== 'undefined') {
                    clearInterval(intervalId);
                    pop.close();
                    this._requestUser(code);
                }

            } catch(err) {
                
            }
        }
        
        let intervalId = setInterval(checkCode, 1000);

    }

    componentWillMount() {
        const avatar = Cookie.get("avatar") || ""
        const name = Cookie.get("name") || ""
        this.setState({
            avatar: avatar,
            name: name
        })
    }


    render() {
        let rightBar = (<div>
            <Icon onClick={this._requestAuthorization} style={{fontSize: 32}} type="github" />
        </div>)
        let disabled = "disabled";
        if(this._isLogged()) {
            disabled = "";
            let btnText = "评论"
            if(this.state.replyer_name !== "") {
                btnText = `回复　${this.state.replyer_name}`
            }
            rightBar = (<div style={{paddingLeft: 12}}>
                <div style={{marginLeft: 8, marginTop: 6}}><img  style={{width: 36, height: 36, borderRadius: 18}}　alt={this.state.name} src={this.state.avatar}></img></div>
                <br/>
                <Button onClick={this._saveComment}>{btnText}</Button>
                </div>)
        }
        const comments = this.state.comments.map((comment, index
            ) => {
            let subcomments = comment.sub_comments.map((subcomment, subindex) => {
                return (<div style={{marginTop: 12, padding: 12}}  key={subindex}>
                <Row>
                    <Col md={{span: 4}}>
                        <img style={{width: 24, height: 24, borderRadius: 12}} src={subcomment.commenter_avatar} />{subcomment.commenter_name}
                    </Col>
                    <Col md={{span: 6}}>
                        {getLocalTime(subcomment.create_time)}
                    </Col>
                    <Col md={{span: 2}}>
                        <Icon onClick={() => this._setComment(comment.comment_id, subcomment.commenter_id, subcomment.commenter_name)} type="message" />
                    </Col>
                    <Col style={{marginTop: 24}} md={{span: 24}}>
                        {subcomment.replyer_id === "" ? subcomment.content : "@" + subcomment.replyer_name + " " + subcomment.content}
                    </Col>
                </Row>
            </div>)
            })
            return (<div style={{marginTop: 36, padding: 12, borderBottom: '1px solid blue'}}  key={index}>
                <Row>
                    <Col md={{span: 4}}>
                        <img style={{width: 24, height: 24, borderRadius: 12}} src={comment.commenter_avatar} />{comment.commenter_name}
                    </Col>
                    <Col md={{span: 6}}>
                       {getLocalTime(comment.create_time)}
                    </Col>
                    <Col md={{span: 2}}>
                        <Icon onClick={() => this._setComment(comment.comment_id, "", comment.commenter_name)} type="message" />
                    </Col>
                    <Col style={{marginTop: 24}} md={{span: 24}}>
                        {comment.content}
                    </Col>
                </Row>
                <div style={{paddingLeft: 48}}>
                    {subcomments}
                </div>
            </div>)
        });
        return (<div>
            <div>
                <Row>
                    <Col md={{span: 18}}>
                        <TextArea value={this.state.comment_behind} onChange={this._handleTextChange}　style={{resize: 'none'}} rows={4} disabled={disabled}>

                        </TextArea>
                    </Col>
                    <Col md={{span: 6}}>
                        {rightBar}
                    </Col>
                </Row>
            </div>
            {comments}
        </div>);
    }
}