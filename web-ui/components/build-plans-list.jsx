import React from "react"

export default function(props) {

    function renderBuildCircles(buildPlans) {
        return Object.keys(buildPlans).map(function (imageName) {
            let buildPlan = buildPlans[imageName];
            let onClick = () => props.onBuildPlanClick(imageName);
            let lastBuild = buildPlan[buildPlan.length - 1];

            let color = lastBuild.exitCode === undefined ?  "lightblue" : (lastBuild.exitCode == 0 ? "lightgreen" : "lightcoral");
            let circleStyle = {
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
