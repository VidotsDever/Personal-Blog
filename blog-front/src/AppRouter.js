import React, { Component } from 'react';
import 'antd/dist/antd.css';
import {BrowserRouter as Router, Route, Link, Switch, Redirect} from 'react-router-dom';
import ForemostPage from './container/Foremost/index';
import EditorPage from './container/Editor/index';

class AppRouter extends Component {
    render() {
        return (
            <Router>
                <Switch>
                    <Route path="/editor" component={EditorPage} />
                    <Route path="/" component={ForemostPage} />
                </Switch>
            </Router>
        );
    }
}

export default AppRouter;