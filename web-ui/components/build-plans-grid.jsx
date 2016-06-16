import React from "react";
import { Link } from "react-router";
import { fromExitCode } from "../css-js/build-colors";
import { shadow, size, cursor, withSiblings } from "../css-js/small-build-squares";

var buildStyle = $.extend(shadow, size, cursor, withSiblings);
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

            const subs = []
            for(let i in sortedBuildPlans) {
                if (i == 0) {
                    continue;
                }

                const image = sortedBuildPlans[i];
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
        const aa = buildPlans[a];
        const bb = buildPlans[b];

        return bb.length - aa.length;
    });
}

function fullWidthBuildPlan(buildPlan) {
    const mostRecentBuild = buildPlan[buildPlan.length - 1];
    const color = fromExitCode(mostRecentBuild.exitCode);
    const styles = {
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
    const mostRecentBuild = buildPlan[buildPlan.length - 1];
    const color = fromExitCode(mostRecentBuild.exitCode);
    const styles = {
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

    const maxBuilds = 8;
    let ellipsing = false;
    if (buildPlan.length > maxBuilds) {
        const startIndex = buildPlan.length - (maxBuilds + 1); // +1 for the ellips
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

    const squares = [];
    if (prefix) {
        squares.push(<div style={baseStyle}>...</div>);
    }

    const lastBuild = builds[builds.length - 1];
    builds.map(build => {
        const color = fromExitCode(build.exitCode);

        const localBuildStyle = JSON.parse(JSON.stringify(buildStyle)); // Hehe.. :D
        localBuildStyle.backgroundColor = color;
        if (build === lastBuild) {
            localBuildStyle.marginRight = "0px";
        }

        const square = <div style={localBuildStyle} title={"Build " + build.id}></div>;
        squares.push(<Link to={"/" + build.id}>{square}</Link>)
    });
    return squares;
}
