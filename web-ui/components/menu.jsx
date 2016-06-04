import React from "react"
import { Link } from "react-router"

export default function(props) {
    let linuxUrl = server + "/client/linux"
    let winUrl = server + "/client/win"

    let styles = {
        float: "right",
        margin: "10px",
        border: "1px solid black",
        padding: "17px",
        paddingBottom: "10px",
        height: "100%",
    };

    return <div style={{display: "flex", flexDirection: "row", flexWrap: "nowrap", justifyContent: "space-between"}}>
            <Link to="/">
                <div style={{float: "left", margin: styles.margin, color: "black"}}>
                    <h1 style={{fontFamily: "jura", margin: "0px"}}>SUAB</h1>
                    <h2 style={{fontFamily: "lavanderia", fontSize: "2.5em", margin: "0px"}}> - we put a shell in your build</h2>
                </div>
            </Link>
            <div></div>
            <fieldset style={styles}>
                <legend>Download client</legend>
                <a href={linuxUrl} style={{marginLeft: "20px"}}><img src='images/linux.png' /></a> &nbsp;
                <a href={winUrl} style={{marginLeft: "20px"}}><img src='images/windows.png' /></a>
            </fieldset>
        </div>
}
