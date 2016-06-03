import React from "react"

export default function(props) {

    function renderBuilds(builds) {
        if (builds.length == 0) {
            return <div>No builds</div>
        }

        let buildStyle = {
            width: "30px",
            height: "30px",
            marginRight: "10px",
            marginTop: "10px",

            WebkitBoxShadow: "1px 1px 5px 1px #555",
            MozBoxShadow: "1px 1px 5px 1px #555",
            BoxShadow: "1px 1px 5x 1px #555",
            borderRadius: "3px",

            cursor: "pointer",

            float: "left",
        };

        let lastBuildId = builds[builds.length - 1].id;
        return builds.map(build => {
            let color = build.exitCode === undefined ?  "lightblue" : (build.exitCode == 0 ? "lightgreen" : "lightcoral");

            let localBuildStyle = JSON.parse(JSON.stringify(buildStyle)); // Hehe.. :D
            localBuildStyle.backgroundColor = color;
            if (build.id === lastBuildId) {
                //localBuildStyle.marginRight = "0px";
            }

            let onClick = () => props.onBuildClick(build.id);

            return <div style={localBuildStyle} onClick={onClick} title={"Build " + build.id}></div>
        });
    }

    let builds = "No builds";
    if (props.builds && Object.keys(props.builds).length > 0) {
        builds = renderBuilds(props.builds);
    }

    return <div>
            {builds}
    </div>
}
