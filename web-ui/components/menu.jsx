window.Menu = function(props) {
    let linuxUrl = server + "/client/linux"
    let winUrl = server + "/client/win"
    return <div>
        <a href={linuxUrl}>download linux client</a> &nbsp;
        <a href={winUrl}>download windows client</a>
    </div>
}
