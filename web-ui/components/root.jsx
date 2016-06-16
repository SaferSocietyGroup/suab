import React from "react"
import Menu from "./menu.jsx"
import { Link } from "react-router"

export default function(props) {


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
