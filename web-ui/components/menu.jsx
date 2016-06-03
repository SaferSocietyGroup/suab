import React from "react"

export default function(props) {
    let linuxUrl = server + "/client/linux"
    let winUrl = server + "/client/win"

    let styles = {
        float: "right",
        paddingRight: "20px",
    };

    return <div style={styles}>
        <a href={linuxUrl}>download linux client</a> &nbsp;
        <a href={winUrl}>download windows client</a>
    </div>
}
