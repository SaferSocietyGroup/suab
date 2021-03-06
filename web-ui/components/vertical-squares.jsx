import React from "react";
import { Link } from "react-router";
import BuildList from "./build-list.jsx";
import Build from "./build.jsx";
import { fromExitCode } from "../css-js/build-colors";

export default function (props) {

    let buildPlans = props.buildPlans;
    if (!buildPlans || Object.keys(buildPlans).length == 0) {
        return <p>n/a</p>;
    }

    const sortedBuildPlans = sortBuildPlans(buildPlans);
    const selectedBuild = props.selectedBuild;
    const selectedBuildPlan = buildPlans[selectedBuild.image];

    const squares = sortedBuildPlans.map(image => renderSquare(buildPlans[image], selectedBuild.image));
    const buildInfo = renderBuildArea(selectedBuildPlan, selectedBuild);

    return <div style={{display: "flex", flexDirection: "row", flexWrap: "nowrap", alignContent: "flex-start"}}>
        <div style={{marginLeft: "10px", maxWidth: "140px"}}>{squares}</div>
        <div style={{marginLeft: "70px"}}>{buildInfo}</div>
    </div>
}

function sortBuildPlans(buildPlans) {
    return Object.keys(buildPlans).sort((a, b) => {
        const aa = buildPlans[a];
        const bb = buildPlans[b];

        return bb.length - aa.length;
    });
}

const shadow = {
    WebkitBoxShadow: "1px 1px 5px 1px #555",
    MozBoxShadow: "1px 1px 5px 1px #555",
    BoxShadow: "1px 1px 5x 1px #555",
    borderRadius: "3px",
};
function renderSquare(buildPlan, selectedImage) {
    const mostRecentBuild = buildPlan[buildPlan.length - 1];
    const color = fromExitCode(mostRecentBuild.exitCode);
    let buildStyle = {
        height: "70px",
        marginRight: "10px",
        marginTop: "10px",
        backgroundColor: color,
        paddingLeft: "1em",
        paddingRight: "1em",
        fontFamily: "jura",
        fontSize: "1.4em",

        cursor: "pointer",

        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        overflow: "hidden",

        color: "black",
    };

    if (buildPlan[0].image === selectedImage) {
        buildStyle = $.extend(buildStyle, shadow);
    }

    return <Link to={"/" + buildPlan[buildPlan.length - 1].id} style={{textDecoration: "none"}}>
        <div style={buildStyle}>{buildPlan[0].image}</div>
    </Link>
}

function renderBuildArea(buildPlan, build) {
    return <div>
            <BuildList builds={buildPlan} selectedBuildId={build.id} />
            <div style={{height: "40px"}}></div>
            <Build build={build} />
    </div>
}
