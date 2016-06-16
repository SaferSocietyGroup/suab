import React from "react"

export default function(props) {

    function renderBuildCircles(buildPlans) {
        return Object.keys(buildPlans).map(function (imageName) {
            const buildPlan = buildPlans[imageName];
            const onClick = () => props.onBuildPlanClick(imageName);
            const lastBuild = buildPlan[buildPlan.length - 1];

            const color = lastBuild.exitCode === undefined ?  "lightblue" : (lastBuild.exitCode == 0 ? "lightgreen" : "lightcoral");
            const circleStyle = {
                borderRadius: "50%",
                width: "200px",
                height: "200px",

                backgroundColor: color,
                display: "flex",
                justifyContent: "center",
                alignItems: "center",

                cursor: "pointer",

                float: "left",
                marginLeft: "10px",
            };

            return <div style={circleStyle} onClick={onClick}>
               {imageName}
            </div>
        });
    }


    let builds = "No builds";
    if (props.buildPlans && Object.keys(props.buildPlans).length > 0) {
        builds = renderBuildCircles(props.buildPlans);
    }

    return <div style={{padding: "20px"}}>
            {builds}
    </div>
}
