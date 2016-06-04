import React from "react"
import { Link } from "react-router"
import {successColor, failColor, unknownColor} from "../css-js/build-colors";

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

            cursor: "pointer",

            float: "left",
        };

        let lastBuildId = builds[builds.length - 1].id;
        return builds.map(build => {
            let color = build.exitCode === undefined ?  unknownColor : (build.exitCode == 0 ? successColor : failColor);

            let localBuildStyle = JSON.parse(JSON.stringify(buildStyle)); // Hehe.. :D
            localBuildStyle.backgroundColor = color;
            if (build.id === lastBuildId) {
                //localBuildStyle.marginRight = "0px";
            }


            return <Link to={"/" + build.id}>
                <div style={localBuildStyle} title={"Build " + build.id}></div>
            </Link>
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
