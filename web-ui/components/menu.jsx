import React from "react"

export default function(props) {
    let linuxUrl = server + "/client/linux"
    let winUrl = server + "/client/win"

    let styles = {
        float: "right",
        margin: "10px",
        border: "1px solid black",
        padding: "17px",
        paddingBottom: "10px",
    };

    return <fieldset style={styles}>
        <legend>Download client</legend>

        <a href={linuxUrl} style={{marginLeft: "20px"}}><img src='images/linux.png' /></a> &nbsp;
        <a href={winUrl} style={{marginLeft: "20px"}}><img src='images/windows.jpg' /></a>
    </fieldset>
}
