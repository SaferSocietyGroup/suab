import React from "react"

export default function(props) {

    function renderBuilds(builds) {
        let buildStyle = {
            width: "30px",
            height: "30px",
            marginRight: "10px",

            float: "left",
        };

        return builds.map(build => {
            let color = build.exitCode === undefined ?  "lightblue" : (build.exitCode == 0 ? "lightgreen" : "lightcoral");

            let localBuildStyle = JSON.parse(JSON.stringify(buildStyle)); // Hehe.. :D
            localBuildStyle.backgroundColor = color;

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
