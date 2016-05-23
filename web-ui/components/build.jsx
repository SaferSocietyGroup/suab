window.Build = React.createClass({

    getInitialState: function() {
        return {
            artifacts: [],
        };
    },

    componentDidMount() {
        $.getJSON("http://localhost:8080/build/" + this.props.buildId + "/artifacts")
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
            artifacts = this.state.artifacts.map(artifact => {
                //let href = "http://localhost:8080/build/" +this.props.buildId+ "/artifacts/" + artifact;
                //return <a href={href}>{artifact}</a>
                return <ArtifactPreview buildId={this.props.buildId} artifactName={artifact} />
            });
        }
        return <div>
            <h2>Build {this.props.buildId}</h2>
            <div>
                <h3>Artifacts</h3>
                {artifacts}
            </div>
        </div>
    }
});
