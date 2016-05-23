window.ArtifactPreview = React.createClass({

    getInitialState: function() {
        return {
            data: "loading...",
        };
    },

    componentWillMount: function() {
        $.get("http://localhost:8080/build/" + this.props.buildId + "/artifacts/" + this.props.artifactName)
            .success(data => {
                console.log(data);
                this.setState({data: data.substring(0, 200)});
            }.bind(this))
            .fail(() => {
                console.log("Could not download artifact", arguments);
                this.setState({data: "Could not download artifact, see console log"});
            }.bind(this));
    },

    render: function() {
        return <div>
            {this.props.artifactName}
            {this.state.data}
        </div>
    }
});