import React from "react"
import Menu from "./menu.jsx"
import { Link } from "react-router"

export default function(props) {

    /*navigateToBuildPlan(buildPlan) {
        console.log("Navigating to build plan", buildPlan, this);
        <BuildList builds={this.state.buildPlans[buildPlan]} />
    }

    navigateToBuild(buildId) {
        let buildPlan = this.findImageFromBuildId(buildId);
        console.log("Navigating to build", buildId, "in build plan", buildPlan);
        <div>
            <BuildList builds={this.state.buildPlans[buildPlan]} />
            <Build buildId={buildId} />
        </div>
    }*/

    let footer /*= <footer style={{ position:"absolute", bottom:"0px", width:"100%", height:"30px", backgroundColor: "#FFF" }}>
            The linux and windows logos made by <a href='http://www.freepik.com/'>Freepic</a> from www.flaticon.com
        </footer>*/;

    return <div>
        <div>
            <Menu />
        </div>
        <div style={{padding: "10px", clear: "both"}}>
            {props.children}
        </div>
        {footer}
    </div>
}
