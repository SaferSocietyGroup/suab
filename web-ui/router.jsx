import React from "react";
import { Router, Route, hashHistory, IndexRoute } from 'react-router';
import App from "./components/root.jsx";
import BuildPlans from "./components/build-plans.jsx";

export default <Router history={hashHistory}>
    <Route path="/" component={App}>
        <IndexRoute component={BuildPlans} />
        <Route path="/:buildid" component={BuildPlans} />
    </Route>
</Router>
