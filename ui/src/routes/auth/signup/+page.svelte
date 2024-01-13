<script lang="ts">
    import {
        UserRegisterRequest,
        validatePassword,
        validateUsername,
        validateEmail,
        EncedeusRegistryApi
    } from "@encedeus/registry-js-api";
    import Input from "$lib/components/Input.svelte";
    import Checkbox from "$lib/components/Checkbox.svelte";
    import Button from "$lib/components/Button.svelte";
    import ErrorTextBox from "$lib/components/ErrorTextBox.svelte";
    import {goto} from "$app/navigation";
    import {accessTokenStore} from "$lib/stores/accessTokenStore";
    import {getApi} from "$lib/api/api";

    let username: string = "", email: string = "", password: string = "";
    let nameError: string = "", emailError: string = "", passwordError: string = "", responseError = "";
    let showPassword: boolean;

    const api = getApi()

    function resetErrors() {
        nameError = "";
        emailError = "";
        passwordError = "";
        responseError = "";
    }

    async function handleSignUp() {

        resetErrors();

        const nameErr = validateUsername(username);
        const emailErr = validateEmail(email);
        const passErr = validatePassword(password);

        if (nameErr) {
            nameError = nameErr.message;
        }
        if (emailErr) {
            emailError = emailErr.message;
        }
        if (passErr) {
            passwordError = passErr.message;
        }

        if (nameErr || emailErr || passErr) {
            return;
        }

        const request: UserRegisterRequest = {
            name: username,
            email: email,
            password: password
        };

        const response = await api.AuthService.SignUp(request);

        if (response.error) {
            responseError = response.error.message;
            return;
        }

        goto("/");
    }

    export function handleShowHidePassword() {
        showPassword = !showPassword;
    }

</script>

<div id="page">

    <Input type="text" placeholder="name" bind:value={username}
           bind:helperText={nameError}/>

    <Input type="email" placeholder="email" bind:value={email}
           bind:helperText={emailError}/>

    <Input type={showPassword ? "text" : "password"} placeholder="password" bind:value={password}
           bind:helperText={passwordError}/>

    <br>

    <Checkbox onclick={handleShowHidePassword}>Show password</Checkbox>

    <br>

    <ErrorTextBox bind:value={responseError}/>

    <Button onclick={handleSignUp}>Sign Up</Button>
</div>

<style>
    #page {
        text-align: center;
    }
</style>
