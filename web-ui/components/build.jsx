window.Build = React.createClass({

    getInitialState: function() {
        return {
            artifacts: [],
        };
    },

    componentDidMount() {
        $.getJSON(server + "/build/" + this.props.buildId + "/artifacts")
            .success(response => {
                this.setState({artifacts: response});
            })
            .fail(() => {
                console.log("Failed loading artifacts", arguments);
                this.setState({artifacts: null});
            });
    },

    render: function() {

        let artifacts;
        if (this.state.artifacts === null) {
            artifacts = "Failed to load artifacts, see console log";
        } else if (this.state.artifacts.length === 0) {
            artifacts = "no artifacts";
        } else {
            <div>
                <h3>Artifacts</h3>
                {artifacts = this.state.artifacts.map(artifact => {
                    return <ArtifactPreview buildId={this.props.buildId} artifactName={artifact} />
                })}
            </div>
        }
        return <div style={{clear: "both"}}>
            <h2>Build {this.props.buildId}</h2>
            {artifacts}
        </div>
    }
});
