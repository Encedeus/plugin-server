<script lang="ts">
    import {
        EncedeusRegistryApi,
        UserSignInRequest,
        validatePassword,
        validateUserIdentifier
    } from "@encedeus/registry-js-api";
    import Input from "$lib/components/Input.svelte";
    import Checkbox from "$lib/components/Checkbox.svelte";
    import Button from "$lib/components/Button.svelte";
    import ErrorTextBox from "$lib/components/ErrorTextBox.svelte";
    import {goto} from "$app/navigation";
    import {getApi} from "$lib/api/api";

    const api = getApi()

    let password: string = "", userIdentifier: string = "";
    let passwordError: string = "", userIdentifierError: string = "", responseError = "";
    let showPassword: boolean;

    async function handleLogin() {

        passwordError = "";
        userIdentifierError = "";
        responseError = "";

        const passErr = validatePassword(password);
        const uidErr = validateUserIdentifier(userIdentifier);

        if (passErr) {
            passwordError = passErr.message;
        }
        if (uidErr) {
            userIdentifierError = uidErr.message;
        }

        if (passErr || uidErr) {
            return;
        }

        const request: UserSignInRequest = {
            uid: userIdentifier,
            password: password
        };
        console.log("sign infsd");
        const response = await api.AuthService.SignIn(request);

        if (response.error) {
            responseError = response.error.message;
            return;
        }

        api.AccessToken = <string>response.response?.accessToken
        console.log("dfwds")
        goto("/")
    }

    export function handleShowHidePassword() {
        showPassword = !showPassword;
    }

</script>

<div id="page">

    <Input type="text" placeholder="username / email" bind:value={userIdentifier}
           bind:helperText={userIdentifierError}/>
    <Input type={showPassword ? "text" : "password"} placeholder="password" bind:value={password}
           bind:helperText={passwordError}/>

    <br>

    <Checkbox onclick={handleShowHidePassword}>Show password</Checkbox>

    <br>

    <ErrorTextBox bind:value={responseError}/>

    <Button onclick={handleLogin}>Login</Button>
</div>

<style>
    #page {
        text-align: center;
    }
</style>
