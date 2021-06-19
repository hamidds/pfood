import axios from "axios";

export default function validateInfo(values) {
    let errors = {}

    // phone number
    if (!values.email.trim()) {
        errors.email = "Please enter email";
    } else {

    }

    // password
    if (!values.password) {
        errors.password = "Please enter password";
    }

    return errors;
}