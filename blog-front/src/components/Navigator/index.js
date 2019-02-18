import React, {Component} from 'react';
import {Menu} from 'antd';
import Cookie from 'js-cookie';
import { Link } from 'react-router-dom';

export default class Navigator extends Component {
    render() {
        const logged = Cookie.get("token") === undefined;
        return (
            <Menu
                defaultSelectedKeys={["1"]}
                mode="horizontal">
                <Menu.Item key="1"><Link to="/">首页</Link></Menu.Item>
                <Menu.Item key="2">收藏</Menu.Item>
                <Menu.Item key="3">日记</Menu.Item>
                <Menu.Item visible={logged} key="4">草稿</Menu.Item>
            </Menu>
        );
    }
}