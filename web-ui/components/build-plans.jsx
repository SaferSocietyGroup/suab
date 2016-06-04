import React from "react"
import BuildPlans from "./build-plans-grid.jsx"
import BuildDetails from "./vertical-squares.jsx"
import { Link } from "react-router"

export default class Root extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            buildPlans: {},
        }
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

    findImageFromBuildId(buildId) {
        for (let image of Object.keys(this.state.buildPlans)) {
            let buildPlan = this.state.buildPlans[image];
            for (let build of buildPlan) {
                if (build.id === buildId) {
                    return image;
                }
            }
        }

        return undefined;
    }

    findBuildById(buildId) {
        for (let image of Object.keys(this.state.buildPlans)) {
            let buildPlan = this.state.buildPlans[image];
            for (let build of buildPlan) {
                if (build.id === buildId) {
                    return build;
                }
            }
        }

        return undefined;
    }

    render() {
        if (this.props.params && this.props.params.buildid) {
            let build = this.findBuildById(this.props.params.buildid);

            return <BuildDetails
                        buildPlans={this.state.buildPlans}
                        selectedBuild={build} />
        } else {
            return <BuildPlans buildPlans={this.state.buildPlans} />
        }
    }
}
