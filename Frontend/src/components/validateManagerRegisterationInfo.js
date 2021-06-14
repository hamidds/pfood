export default function validateInfo(values) {
    let errors = {}

    // email
    if (!values.email.trim()) {
        errors.email = "Email required";
    }

    // password
    if (!values.password) {
        errors.password = "Password is required";
    }
    else if (values.password.length < 8) {
        errors.password = "Password needs to be 8 character or more";
    }

    return errors;
}