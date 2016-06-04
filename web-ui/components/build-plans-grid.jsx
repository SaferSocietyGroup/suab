import React from "react";
import { Link } from "react-router";
import {successColor, failColor, unknownColor} from "../css-js/build-colors";

export default function(props) {
    let buildPlans = props.buildPlans;
    if (!buildPlans || Object.keys(buildPlans).length == 0) {
        return <p>No builds</p>
    }

    // sort plans by number of builds in them
    // If there is one that has more than all other, give it the whole width
    // All others equal width,

    let sortedBuildPlans = sortBuildPlans(buildPlans);

    if (sortedBuildPlans.length > 1) {
        let p = [];
        if (buildPlans[sortedBuildPlans[0]].length > buildPlans[sortedBuildPlans[1]].length) {

            p.push(fullWidthBuildPlan(buildPlans[sortedBuildPlans[0]]));

            let subs = []
            for(let i in sortedBuildPlans) {
                if (i == 0) {
                    continue;
                }

                let image = sortedBuildPlans[i];
                subs.push(notFullWidth(buildPlans[image]));
            }

            p.push(<div style={{marginLeft: "-5px"}}>{subs}</div>);
        } else {
            for(let image in sortedBuildPlans) {
                p.push(notFullWidth(buildPlans[image]));
            }
        }

        return <div>{p}</div>
    } else {
        return fullWidthBuildPlan(buildPlans[sortedBuildPlans[0]]);
    }
}


function sortBuildPlans(buildPlans) {
    return Object.keys(buildPlans).sort((a, b) => {
        let aa = buildPlans[a];
        let bb = buildPlans[b];

        return bb.length - aa.length;
    });
}

function fullWidthBuildPlan(buildPlan) {
    let mostRecentBuild = buildPlan[buildPlan.length - 1];
    let color = mostRecentBuild.exitCode === undefined ? unknownColor : (mostRecentBuild.exitCode == 0 ? successColor : failColor);
    let styles = {
        width: "100%",
        height: "350px",
        marginBottom: "5px",
        backgroundColor: color,
        fontFamily: "jura",

        display: "flex",
        justifyContent: "center",
        alignItems: "center",
    };

    return <div style={styles}>
        <div>
            <h1 style={{textAlign: "center"}}>{buildPlan[0].image}</h1>
            {renderBuilds(buildPlan)}
        </div>
    </div>
}

function notFullWidth(buildPlan) {
    let mostRecentBuild = buildPlan[buildPlan.length - 1];
    let color = mostRecentBuild.exitCode === undefined ? unknownColor : (mostRecentBuild.exitCode == 0 ? successColor : failColor);
    let styles = {
        width: "24.28vw",
        float: "left",
        marginLeft: "5px",
        marginBottom: "5px",

        height: "200px",
        backgroundColor: color,
        fontFamily: "jura",

        display: "flex",
        justifyContent: "center",
        alignItems: "center",
    };

    let maxBuilds = 8;
    let ellipsing = false;
    if (buildPlan.length > maxBuilds) {
        let startIndex = buildPlan.length - (maxBuilds + 1); // +1 for the ellips
        buildPlan = buildPlan.slice(startIndex, startIndex + maxBuilds);
        ellipsing = true;
    }

    return <div style={styles}>
        <div>
            <h1 style={{textAlign: "center"}}>{buildPlan[0].image}</h1>
            {renderBuilds(buildPlan, ellipsing ? "..." : undefined)}
        </div>
    </div>
}

function renderBuilds(builds, prefix) {
    if (builds.length == 0) {
        return <div>No builds</div>
    }

    let baseStyle = {
        width: "30px",
        height: "30px",
        marginTop: "10px",

        float: "left",
    }

    let buildStyle = $.extend({
        WebkitBoxShadow: "1px 1px 5px 1px #555",
        MozBoxShadow: "1px 1px 5px 1px #555",
        BoxShadow: "1px 1px 5x 1px #555",
        borderRadius: "3px",

        marginRight: "10px",

        cursor: "pointer",
    }, baseStyle);

    let squares = [];
    if (prefix) {
        squares.push(<div style={baseStyle}>...</div>);
    }

    let lastBuild = builds[builds.length - 1];
    builds.map(build => {
        let color = build.exitCode === undefined ?  "lightblue" : (build.exitCode == 0 ? "lightgreen" : "lightcoral");

        let localBuildStyle = JSON.parse(JSON.stringify(buildStyle)); // Hehe.. :D
        localBuildStyle.backgroundColor = color;
        if (build === lastBuild) {
            localBuildStyle.marginRight = "0px";
        }

        let square = <div style={localBuildStyle} title={"Build " + build.id}></div>;
        squares.push(<Link to={"/" + build.id}>{square}</Link>)
    });
    return squares;
}
