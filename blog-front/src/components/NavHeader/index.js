import React, { Component } from 'react';
import './index.css';
import AuthBar from '../AuthBar/index';
import Navigator from '../Navigator/index';
import { Layout, Row, Col } from 'antd';
const { Header } = Layout;


export default class NavHeader extends Component {
    render() {
        return (
            <Header className="header-container">
                <Row>
                    <Col md={{span: 4}}></Col>
                    <Col md={{span: 14}}><Navigator></Navigator></Col>
                    <Col md={{span: 6}}><AuthBar></AuthBar></Col>
                </Row>
            </Header>
        );
    }
}