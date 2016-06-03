import React from "react"
import Menu from "./menu.jsx"
import BuildList from "./build-list.jsx"
import BuildPlansList from "./build-plans-list.jsx"
import Build from "./build.jsx"

export default class Root extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            buildPlans: [],
            content: null,
        };
    }

    componentDidMount() {
        $.getJSON(server + '/builds')
            .success(function(response) {
                if (!response) {
                    response = [];
                } else {
                    let arr = [];
                    for (var key in response) {
                        let obj = {};
                        obj[key] = response[key];
                        arr.push(obj);
                    }
                    response = arr;
                }
                let buildPlans = this.groupBuildsByImage(response);
        	    this.setState({buildPlans: buildPlans});
            }.bind(this))
            .error(function(err) {
        	    console.log("ERROR", err, this);
        	    this.setState({buildPlans: null});
            }.bind(this));
    }

    groupBuildsByImage(builds) {
        var images = {};
        for (let i in builds) {
            let buildId = Object.keys(builds[i])[0];
            let build = builds[i][buildId];
            let key = build.image;

            if (!images[key]) {
                images[key] = [];
            }

            build.id = buildId;
            images[key].push(build);
        }

        return images;
    }

    navigateToBuildPlan(buildPlan) {
        console.log("Navigating to build plan", buildPlan, this);
        this.setState({content: <BuildList builds={this.state.buildPlans[buildPlan]} onBuildClick={this.navigateToBuild.bind(this)} />});
    }

    navigateToBuild(buildId) {
        let buildPlan = Object.keys(this.state.buildPlans).filter(planKey => this.state.buildPlans[planKey].filter(build => build.id === buildId))[0];
        console.log("Navigating to build", buildId, "in build plan", buildPlan);
        this.setState({content: <div>
            <BuildList builds={this.state.buildPlans[buildPlan]} onBuildClick={this.navigateToBuild.bind(this)} />
            <Build buildId={buildId} />
        </div>});
    }

    render() {

        var buildList = undefined;
        if (this.state.buildPlans === null) {
            buildList = "Could not load builds, see console log";
        } else {
            buildList = <BuildPlansList buildPlans={this.state.buildPlans} onBuildPlanClick={this.navigateToBuildPlan.bind(this)}/>
        }

        return <div>
            <div>
                <Menu />
            </div>
            <div>
                {buildList}
            </div>
            <div style={{padding: "10px", clear: "both"}}>
                {this.state.content}
            </div>
            <footer style={{ position:"absolute", bottom:"0", width:"100%", height:"30px", backgroundColor: "#FFF" }}>
                The linux and windows logos made by <a href='http://www.freepik.com/'>Freepic</a> from www.flaticon.com
            </footer>
        </div>
    }
}
