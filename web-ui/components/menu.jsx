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

    return <div>
            <div style={{float: "left", margin: styles.margin}}>
                <h1 style={{fontFamily: "jura", margin: "0px"}}>SUAB</h1>
                <h2 style={{fontFamily: "lavanderia", fontSize: "2.5em", margin: "0px"}}> - we put a shell in your build</h2>
            </div>

            <fieldset style={styles}>
                <legend>Download client</legend>
                <a href={linuxUrl} style={{marginLeft: "20px"}}><img src='images/linux.png' /></a> &nbsp;
                <a href={winUrl} style={{marginLeft: "20px"}}><img src='images/windows.jpg' /></a>
            </fieldset>
        </div>
}
