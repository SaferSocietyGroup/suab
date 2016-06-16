import React from "react"
import { Link } from "react-router"

const linuxUrl = server + "/client/linux"
const winUrl = server + "/client/win"

const menuContainerStyle = {display: "flex", flexDirection: "row", flexWrap: "nowrap", justifyContent: "space-between"};

const downloadContainerStyle = {
    float: "right",
    margin: "10px",
    border: "1px solid black",
    padding: "17px",
    paddingBottom: "10px",
    height: "100%",
};
const downloadLinkStyle = {marginLeft: "20px"};

const headerContainerStyle ={float: "left", margin: downloadContainerStyle.margin, color: "black"};
const headerStyle = {fontFamily: "jura", margin: "0px"};
const tagLineStyle = {fontFamily: "lavanderia", fontSize: "2.5em", margin: "0px"};

export default function(props) {

    return <div style={menuContainerStyle}>
            <Link to="/">
                <div style={headerContainerStyle}>
                    <h1 style={headerStyle}>SUAB</h1>
                    <h2 style={tagLineStyle}> - we put a shell in your build</h2>
                </div>
            </Link>
            <div></div>
            <fieldset style={downloadContainerStyle}>
                <legend>Download client</legend>
                <a href={linuxUrl} style={downloadLinkStyle}><img src='images/linux.png' /></a> &nbsp;
                <a href={winUrl} style={downloadLinkStyle}><img src='images/windows.png' /></a>
            </fieldset>
        </div>
}
