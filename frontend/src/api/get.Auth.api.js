import {getRequest} from "./RequestUtilities";

export function getAuth() {
    return getRequest("/auth/account" );
}