import React from "react"
import BuildPlans from "./build-plans-grid.jsx"
import BuildDetails from "./vertical-squares.jsx"
import { Link } from "react-router"

export default class Root extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            loadingState: "not-loaded",
            lastError: null,
            buildPlans: {},
        }
    }

    componentDidMount() {
        $.getJSON(server + '/builds')
            .success(function(response) {
                if (!response) {
                    response = [];
                } else {
                    const arr = [];
                    for (var key in response) {
                        const obj = {};
                        obj[key] = response[key];
                        arr.push(obj);
                    }
                    response = arr;
                }
                const buildPlans = this.groupBuildsByImage(response);
                this.setState({
                    loadingState: "loaded",
                    lastError: null,
                    buildPlans: buildPlans,
                });
            }.bind(this))
            .error(function(err, b, c) {
                this.setState({
                    loadingState: "error",
                    lastError: "Failed loading builds from server, please check the network tab in the dev tools :)",
                    buildPlans: null
                });
            }.bind(this));
    }

    groupBuildsByImage(builds) {
        const images = {};
        for (let i in builds) {
            const buildId = Object.keys(builds[i])[0];
            const build = builds[i][buildId];
            const key = build.image;

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
            const buildPlan = this.state.buildPlans[image];
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
            const buildPlan = this.state.buildPlans[image];
            for (let build of buildPlan) {
                if (build.id === buildId) {
                    return build;
                }
            }
        }

        return undefined;
    }

    render() {
        if (this.state.loadingState === "loaded") {
            return this.renderBuilds();
        } else if (this.state.loadingState === "error") {
            return <div>{this.state.lastError}</div>
        } else {
            return <div>Loading...</div>
        }
    }

    renderBuilds() {
        if (this.props.params && this.props.params.buildid) {
            const build = this.findBuildById(this.props.params.buildid);

            return <BuildDetails
                        buildPlans={this.state.buildPlans}
                        selectedBuild={build} />
        } else {
            return <BuildPlans buildPlans={this.state.buildPlans} />
        }
    }
}
