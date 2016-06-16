import React from "react"
import { Link } from "react-router"
import { fromExitCode } from "../css-js/build-colors";
import { size, cursor, withSiblings } from "../css-js/small-build-squares";

const buildStyle = $.extend(size, cursor, withSiblings, { marginTop: "10px"});

export default function(props) {
    let builds = "No builds";
    if (props.builds && Object.keys(props.builds).length > 0) {
        builds = renderBuilds(props.builds);
    }

    return <div>
            {builds}
    </div>
}

function renderBuilds(builds) {
    if (builds.length == 0) {
        return <div>No builds</div>
    }

    const lastBuildId = builds[builds.length - 1].id;
    return builds.map(build => {
        const color = fromExitCode(build.exitCode);

        const localBuildStyle = JSON.parse(JSON.stringify(buildStyle)); // Hehe.. :D
        localBuildStyle.backgroundColor = color;

        return <Link to={"/" + build.id}>
            <div style={localBuildStyle} title={"Build " + build.id}></div>
        </Link>
    });
}
