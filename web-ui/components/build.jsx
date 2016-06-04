import React from "react"
import ArtifactPreview from "./artifact-preview.jsx"

export default class Build extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            artifacts: [],
        };
    }

    componentDidMount() {
        $.getJSON(server + "/build/" + this.props.build.id + "/artifacts")
            .success(response => {
                this.setState({artifacts: response});
            })
            .fail(() => {
                console.log("Failed loading artifacts", arguments);
                this.setState({artifacts: null});
            });
    }

    render() {

        let artifacts;
        if (this.state.artifacts === null) {
            artifacts = "Failed to load artifacts, see console log";
        } else if (this.state.artifacts.length === 0) {
            artifacts = "no artifacts";
        } else {
            <div>
                <h3>Artifacts</h3>
                {artifacts = this.state.artifacts.map(artifact => {
                    return <ArtifactPreview buildId={this.props.build.id} artifactName={artifact} />
                })}
            </div>
        }
        return <div style={{clear: "both"}}>
            <h2>Build {this.props.build.id} of {this.props.build.image}</h2>
            {artifacts}
        </div>
    }
}
