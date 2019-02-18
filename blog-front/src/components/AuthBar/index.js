import React, {Component} from 'react';
import {Menu, Button, Dropdown, Icon} from 'antd';
import {Link} from 'react-router-dom';
import Cookie from 'js-cookie';

export default class AuthBar extends Component {
    render() {
        if(Cookie.get("token") !== undefined) {
            return (
                <div style={{float: "right", right: 12}}>
                    <Button
                        ghost
                        type="primary"
                        size="small">
                        登录
                    </Button>
                </div>
            );
        }
        const menu = (
            <Menu>
                <Menu.Item key="0">
                    <Link to="/editor" >写博客</Link>
                </Menu.Item>
                <Menu.Item key="１">
                    <a href="http://www.baidu.com">写日记</a>
                </Menu.Item>
                <Menu.Item key="２">
                    <a href="http://www.baidu.com">添加收藏</a>
                </Menu.Item>
                <Menu.Item key="３">
                    <a href="http://www.baidu.com">退出账户</a>
                </Menu.Item>
            </Menu>
        );
        return (
            <div style={{float: "right", right: 12}}>
                <Dropdown overlay={menu} trigger={['click']}>
                    <a className="ant-dropdown-link" href="#">
                        <Icon type="setting" />
                    </a>
                </Dropdown>
            </div>
        );
    }
}