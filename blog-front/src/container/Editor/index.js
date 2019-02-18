import React, {Component} from 'react';
import GenerateUUID from '../../utils/rand';
import BraftEditor from 'braft-editor';
import 'braft-editor/dist/index.css';
import { Layout, Row, Col, Button, Input } from 'antd';
import axios from 'axios';
import qs from 'qs';
const {Header, Content} = Layout;



export default class EditorPage extends Component {
    
    constructor(props) {
        super(props)
        this.state = {
            identity: "",
            title: "",
            editorState: null
        };
    }

    componentWillMount() { 
        this.setState({
            identity: GenerateUUID()
        })
    }

    _mediaValidateFn = mediaObj => {
        if(mediaObj.size > 1024 * 1024) {
            alert("上传的图片不能超过1M")
        }
        return mediaObj.size < 1024 * 1024
    }

    _handleEditorChange = (editorState) => {
        this.setState({editorState})
    }
    // true表示发布，false表示保存
    _submitContent = async (status) => {
        const htmlContent = this.state.editorState.toHTML()
        const rawContent = this.state.editorState.toRAW()
        const data = {
            id: this.state.identity,
            title: this.state.title,
            html: htmlContent,
            raw: rawContent,
            status: status === true ? "publish" : "save"
        };
        axios.post(
            "http://localhost:8080/media/editor",
            qs.stringify(data),
            {
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                  }
            }
        ).then(response => {
            if(response.statusText === "OK") {
                this.props.history.push("/")
            } else {
                console.log(response)
            }
        }).catch(err => {
            console.log(err)
        });
    }

    _uploadFn = param => {
        const serverURL = "http://localhost:8080/media/upload"
        const xhr = new XMLHttpRequest
        const fd = new FormData()

        const successFn = response => {
            param.success({
                url: JSON.parse(xhr.responseText).url,
                meta: {
                    id: 'some-body',
                    title: 'fuck-the-author',
                    alt: xhr.responseText.url,
                    loop: true, // 指定音视频是否循环播放
                    autoPlay: true, // 指定音视频是否自动播放
                    controls: true, // 指定音视频是否显示控制栏
                    poster: xhr.responseText.url, // 指定视频播放器的封面
                }
            })
        }

        const progressFn = event => {
            param.progress(event.loaded / event.total * 100)
        }

        const errorFn = response => {
            param.error({
                msg: '不能上传'
            })
        }

        xhr.upload.addEventListener("progress", progressFn, false)
        xhr.addEventListener("load", successFn, false)
        xhr.addEventListener("error", errorFn, false)
        xhr.addEventListener("abort", errorFn, false)

        fd.append("image", param.file)
        xhr.open("POST", serverURL, true)
        xhr.send(fd)
    }

    _handleTextChange = e => {
        this.setState({
            title: e.target.value
        });
    }
    
    render() {
        const {editorState} = this.state
        return (
            <div>
                <Layout>
                    <Header>
                    <Row>
                        <Col key="1" md={{span: 18}} xs={{span: 12}}>
                            <div>
                                <Input onChange={this._handleTextChange} placeholder="输入标题" />
                            </div>
                        </Col>
                        <Col key="2" md={{span: 6}} xs={{span: 12}}>
                            <Button
                                ghost
                                type="primary"
                                size="small"
                                onClick={ () => this._submitContent(false)}
                                style={{marginRight: 20, marginLeft: 20}}>
                                    保存
                            </Button>
                            <Button
                                ghost
                                type="primary"
                                size="small"
                                onClick={ () => this._submitContent(true)}>
                                    发布
                            </Button>
                        </Col>
                    </Row>
                    </Header>
                    <Content>
                        <BraftEditor 
                            value={editorState}
                            media={{
                                validateFn: this._mediaValidateFn,
                                uploadFn: this._uploadFn,
                                accepts: {
                                    image: 'image/png,image/jpeg,image/gif,image/webp,image/apng,image/svg',
                                    video: false,
                                    audio: false,
                                }
                            }}
                            onChange={this._handleEditorChange}
                            onSave={() => this._submitContent(false)}/>
                    </Content>
                </Layout>
            </div>
        );
    }
}