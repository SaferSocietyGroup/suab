import React from "react"

export default class ArtifactPreview extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            data: "loading...",
        };
    }

    componentWillMount() {
        $.get(server + "/build/" + this.props.buildId + "/artifacts/" + this.props.artifactName)
            .success(this.onFetchSuccess.bind(this))
            .fail(this.onFetchError.bind(this));
    }

    onFetchSuccess(data, _, jqxhr) {
        if (jqxhr.getResponseHeader("Content-Type").indexOf("text") >= 0) {
            this.setState({data: data.substring(0, 200)});
        } else {
            this.setState({data: "Not plain text"});
        }
    }

    onFetchError() {
        console.log("Could not download artifact", arguments);
        this.setState({data: "Could not download artifact, see console log"});
    }

    render() {
        let href = server + "/build/" +this.props.buildId+ "/artifacts/" + this.props.artifactName;

        return <div>
            <a href={href} target="_blank">{this.props.artifactName}</a>
            <br />
            {this.state.data}
        </div>
    }
}
