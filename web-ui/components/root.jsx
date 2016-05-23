window.Root = React.createClass({

    getInitialState: function() {
        return {
            builds: [],
            content: "Choose a build to the left",
        };
    },

    componentDidMount: function() {
        $.getJSON('http://localhost:8080/builds')
            .success(function(response) {
        	    this.setState({builds: Object.keys(response)});
            }.bind(this))
            .error(function(err) {
        	    console.log("ERROR", err, this);
        	    this.setState({builds: null});
            }.bind(this));
    },

    navigateToBuild: function(buildId) {
        console.log("Navigating to", buildId);
        this.setState({content: <Build buildId={buildId}/>});
    },

    render: function() {

        var menu = undefined;
        if (this.state.builds === null) {
            menu = "Could not load builds, see console log";
        } else {
            menu = <Menu builds={this.state.builds} onBuildClick={this.navigateToBuild}/>
        }

        return <div>
            <div>
                {menu}
            </div>
            <div style={{padding: "10px"}}>
                {this.state.content}
            </div>
        </div>
    },
});