export default function validateInfo(values) {
    let errors = {}

    // phone number
    if (!values.phone_number.trim()) {
        errors.phone_number = "Phone number required";
    }
    // password
    if (!values.password) {
        errors.password = "Password is required";
    }

    return errors;
}

