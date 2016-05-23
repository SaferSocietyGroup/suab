window.Menu = function(props) {
    var builds = props.builds.map(function (build) {
        let href = build + "/apa"
        let onClick = () => props.onBuildClick(build)

        let circleStyle = {
            borderRadius: "50%",
            width: "200px",
            height: "200px",

            backgroundColor: "lightgreen",
            display: "flex",
            justifyContent: "center",
            alignItems: "center",

            cursor: "pointer",
        };

        return <div style={circleStyle} onClick={onClick}>
           {build}
        </div>
    });

    return <div style={{padding: "20px"}}>
            {builds}
    </div>
}