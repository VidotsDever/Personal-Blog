import React, { Component } from 'react';
import NavHeader from '../../components/NavHeader'
import { Switch, Route, Redire, Redirect } from 'react-router-dom';
import {Layout} from 'antd';
import PostList from '../../components/PostList/index';
import PostDetail from '../../components/PostDetail/index';
const {Content, Footer } = Layout;


class ForemostPage extends Component {
    render() {
        
        const minHeight = document.body.clientHeight;
        return (
            <div>
                <Layout className="layout">
                    <NavHeader></NavHeader>
                    <Content style={{ padding: '20px 60px' }}>
                        <div style={{ background: '#fff', padding: 24, minHeight: minHeight - 64 - 62 }}>
                            <Switch>
                                <Route path="/" exact render={() => <Redirect to="/posts"></Redirect>} />
                                <Route path="/posts" exact component={PostList} />
                                <Route path="/posts/:id" component={PostDetail} />
                            </Switch>
                        </div>
                    </Content>
                    <Footer style={{ textAlign: 'center' }}>
                        Personal Blog Created By Vidots
                    </Footer>
                </Layout>
            </div>
        );
    }
}

export default ForemostPage;