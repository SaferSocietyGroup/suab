window.ArtifactPreview = React.createClass({

    getInitialState: function() {
        return {
            data: "loading...",
        };
    },

    componentWillMount: function() {
        $.get(server + "/build/" + this.props.buildId + "/artifacts/" + this.props.artifactName)
            .success((data, _, jqxhr) => {
                if (jqxhr.getResponseHeader("Content-Type").indexOf("text") >= 0) {
                    this.setState({data: data.substring(0, 200)});
                } else {
                    this.setState({data: "Not plain text"});
                }
            }.bind(this))
            .fail(() => {
                console.log("Could not download artifact", arguments);
                this.setState({data: "Could not download artifact, see console log"});
            }.bind(this));
    },

    render: function() {
        let href = server + "/build/" +this.props.buildId+ "/artifacts/" + this.props.artifactName;

        return <div>
            <a href={href} target="_blank">{this.props.artifactName}</a>
            <br />
            {this.state.data}
        </div>
    }
});
