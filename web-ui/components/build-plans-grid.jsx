import React from "react";
import BuildList from "./build-list.jsx";

export default function(props) {

    let successColor = "#afa";
    let failColor = "lightsalmon";
    let unknownColor = "lightblue";

    function sortBuildPlans() {
        return Object.keys(props.buildPlans).sort((a, b) => {
            let aa = props.buildPlans[a];
            let bb = props.buildPlans[b];

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

            display: "flex",
            justifyContent: "center",
            alignItems: "center",
        };

        return <div style={styles}>
            <div>
                <h1 style={{textAlign: "center"}}>{buildPlan[0].image}</h1>
                <BuildList builds={buildPlan} onBuildClick={props.onBuildClick}/>
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

            let onClick = () => props.onBuildClick(build.id);

            squares.push(<div style={localBuildStyle} onClick={onClick} title={"Build " + build.id}></div>);
        });
        return squares;
    }

    if (!props.buildPlans || Object.keys(props.buildPlans).length == 0) {
        return <p>No builds</p>
    }

    // sort plans by number of builds in them
    // If there is one that has more than all other, give it the whole width
    // All others equal width,

    let sortedBuildPlans = sortBuildPlans();

    if (sortedBuildPlans.length > 1) {
        let p = [];
        if (props.buildPlans[sortedBuildPlans[0]].length > props.buildPlans[sortedBuildPlans[1]].length) {

            p.push(fullWidthBuildPlan(props.buildPlans[sortedBuildPlans[0]]));

            let subs = []
            for(let i in sortedBuildPlans) {
                if (i == 0) {
                    continue;
                }

                let image = sortedBuildPlans[i];
                subs.push(notFullWidth(props.buildPlans[image]));
            }

            p.push(<div style={{marginLeft: "-5px"}}>{subs}</div>);
        } else {
            for(let image in sortedBuildPlans) {
                p.push(notFullWidth(props.buildPlans[image]));
            }
        }

        return <div>{p}</div>
    } else {
        return fullWidthBuildPlan(props.buildPlans[sortedBuildPlans[0]]);
    }
}
