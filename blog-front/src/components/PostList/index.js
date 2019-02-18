import React, { Component } from 'react';
import {Pagination} from 'antd';
import axios from 'axios';
import { Route, Link } from 'react-router-dom';
import {getLocalTime} from '../../utils/date'

export default class PostList extends Component {

    constructor(props) {
        super(props);
        this.state = {
            posts: [],
            count: 0,
            pagesize: 4
        };
    }

    _fetchPosts = async (page) => {
        axios.get(`http://localhost:8080/blog/posts?page=${page}&pagesize=${this.state.pagesize}`).then(response => {
            this.setState({
                count: response.data.count,
                posts: response.data.posts
            });
        }).catch(err => {
            console.log(err)
        })
    }

    componentDidMount() {
        this._fetchPosts(1)
    }

    _onChangePage = page => {
        this._fetchPosts(page)
    }

    render() {
        if(this.state.count == 0) {
            return (
                <div style={{textAlign: 'center', fontSize: 24}}>
                    :)
                </div>
            );
        }
        const count = this.state.count
        const pagesize = this.state.pagesize
        const posts = this.state.posts.map((post, index) => {
            return (
                <div style={{marginBottom: 36}} key={index}>
                    <h2><Link to={`${this.props.match.url}/${post.id}`}>{post.title}</Link></h2>
                    <strong>{'   '}{getLocalTime(post.create_time)}</strong>
                    <p>{`${post.html_str.replace(/<[^>]*>|/g,"").substring(0, 100)}`}
                        {' '}<Link to={`${this.props.match.url}/${post.id}`}>查看更多</Link>
                    </p>
                </div>
            );
        })
        return (
            <div style={{paddingLeft: 24, paddingRight: 24}}>
                <div>
                {posts}
                </div>
                <Pagination defaultPageSize={pagesize} defaultCurrent={5} total={count} onChange={this._onChangePage} />
            </div>
        );
    }
}