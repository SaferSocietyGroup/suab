window.Root = React.createClass({

    getInitialState: function() {
        return {
            buildPlans: [],
            content: null,
        };
    },

    componentDidMount: function() {
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
    },

    groupBuildsByImage: function (builds) {
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
    },

    navigateToBuildPlan: function(buildPlan) {
        console.log("Navigating to", buildPlan);
        this.setState({content: <BuildList builds={this.state.buildPlans[buildPlan]} onBuildClick={this.navigateToBuild} />});
    },

    navigateToBuild: function(buildId) {
        let buildPlan = Object.keys(this.state.buildPlans).filter(planKey => this.state.buildPlans[planKey].filter(build => build.id === buildId))[0];
        console.log("Navigating to", buildId, buildPlan);
        this.setState({content: <div>
            <BuildList builds={this.state.buildPlans[buildPlan]} onBuildClick={this.navigateToBuild} />
            <Build buildId={buildId} />
        </div>});
    },

    render: function() {

        var buildList = undefined;
        if (this.state.buildPlans === null) {
            buildList = "Could not load builds, see console log";
        } else {
            buildList = <BuildPlansListCircles buildPlans={this.state.buildPlans} onBuildPlanClick={this.navigateToBuildPlan}/>
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
        </div>
    },
});
