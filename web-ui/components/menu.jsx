window.Menu = function(props) {
    console.log(props)
    var builds = props.builds.map(function (build) {
        let href = build + "/apa"
        let onClick = () => props.onBuildClick(build)
        return <a href="javascript:void(0)" onClick={onClick}>{build}</a>
    });

    return <div>
        <h2>Builds</h2>
        <div>
            {builds}
        </div>
    </div>
}