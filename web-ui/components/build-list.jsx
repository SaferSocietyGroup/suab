import React from "react"
import { Link } from "react-router"
import { fromExitCode } from "../css-js/build-colors";
import { size, cursor, withSiblings } from "../css-js/small-build-squares";

const buildStyle = $.extend(size, cursor, withSiblings, { marginTop: "10px"});
const buildContainerStyle = {display: "inline-block"};
const arrowStyle = {position: "absolute", paddingLeft: "10px", marginTop: "-17px", transform: "rotateZ(-30deg)"};

export default function(props) {
    let builds = "No builds";
    if (props.builds && Object.keys(props.builds).length > 0) {
        builds = renderBuilds(props.builds, props.selectedBuildId);
    }

    return <div>
        {builds}
    </div>
}

function renderBuilds(builds, selectedBuildId) {
    if (builds.length == 0) {
        return <div>No builds</div>
    }

    const lastBuildId = builds[builds.length - 1].id;
    return builds.map(build => {
        const color = fromExitCode(build.exitCode);

        const localBuildStyle = JSON.parse(JSON.stringify(buildStyle)); // Hehe.. :D
        localBuildStyle.backgroundColor = color;

        let arrow = undefined;
        if (build.id === selectedBuildId) {
            arrow = <img src="images/arrow.png" style={arrowStyle}/>
        }

        return <div style={buildContainerStyle}>
            {arrow}
            <Link to={"/" + build.id}>
                <div style={localBuildStyle} title={"Build " + build.id}></div>
            </Link>
        </div>
    });
}
