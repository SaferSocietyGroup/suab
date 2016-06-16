export var successColor = "#afa";
export var failColor = "lightsalmon";
export var unknownColor = "lightblue";

export function fromExitCode(exitCode) {
    if (exitCode === undefined) {
       return unknownColor;
    } else if (exitCode == 0) {
       return successColor;
    } else {
       return failColor;
    }
}
