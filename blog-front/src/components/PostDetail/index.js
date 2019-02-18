import React, { Component } from 'react';
import axios from 'axios';
import Comments from '../Comments';

export default class PostDetail extends Component {

    constructor(props) {
        super(props)
        this.state = {
            post: null
        }
    }



    componentDidMount() {
        this._fetchPostByID(this.props.match.params.id)
    }


    _fetchPostByID = async id => {
        axios.get(`http://localhost:8080/blog/post?id=${id}`).then(response => {
            const post = response.data.post
            this.setState({
                post: post,
            });
        }).catch(err => {
            console.log(err)
        })
    }

    render() {
        if(this.state.post === null) {
            return (
                <div>
                    :)
                </div>
            );
        }

        return (
            <div>
                <h2>{this.state.post.title}</h2>
                <div dangerouslySetInnerHTML={{__html: this.state.post.html_str}}></div>
                <Comments post_id={this.props.match.params.id} />
            </div>
        );
    }
}