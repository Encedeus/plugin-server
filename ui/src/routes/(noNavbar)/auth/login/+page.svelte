<script lang="ts">
    import {
        UserSignInRequest,
        validatePassword,
        validateUserIdentifier
    } from "@encedeus/registry-js-api";
    import Input from "$lib/components/generic/Input.svelte";
    import Checkbox from "$lib/components/Checkbox.svelte";
    import Button from "$lib/components/generic/Button.svelte";
    import ErrorTextBox from "$lib/components/ErrorTextBox.svelte";
    import {goto} from "$app/navigation";
    import {getApi} from "$lib/api/api";
    import Card from "$lib/components/generic/Card.svelte";
    import CardHeader from "$lib/components/generic/CardHeader.svelte";
    import SmallArrowRight from "$lib/components/heroicons/SmallArrowRight.svelte";

    const api = getApi();

    let password: string = "", userIdentifier: string = "";
    let passwordError: string = "", userIdentifierError: string = "", responseError = "";
    let showPassword: boolean;

    async function handleLogin() {

        clearError();

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

        const response = await api.AuthService.SignIn(request);

        if (response.error) {
            responseError = response.error.message;
            return;
        }

        //api.AccessToken = <string>response.response?.accessToken
        goto("/");
    }

    function clearError() {
        passwordError = "";
        userIdentifierError = "";
        responseError = "";
    }

    export function handleShowHidePassword() {
        showPassword = !showPassword;
    }

</script>
<!--
<div id="page" class="w-min m-auto h-full">
    <Card className="w-min" width="lg">
        <div slot="content" class="flex flex-col justify-stretch gap-5 mt-2.5">
            <Input label="username / email"
                   type="text"
                   placeholder="username / email"
                   bind:value={userIdentifier}
                   error={!!userIdentifierError}
            />
            <Input label="password"
                   type={showPassword ? "text": "password"}
                   placeholder="password"
                   bind:value={password}
                   error={!!passwordError}
            />

            <Checkbox className="m-auto" onclick={handleShowHidePassword}>Show password</Checkbox>

            <Button className="m-auto" onClick={handleLogin}>Login</Button>

            <ErrorTextBox className=" m-auto" bind:value={responseError}/>
        </div>
    </Card>
</div>
-->

<div class="overflow-hidden">
    <aside class="absolute top-0 right-0 mt-5 mr-7">
        <span class="drop-shadow-xl text-white text-sm font-bold tracking-wide">Don't have an account?&nbsp; • &nbsp;<a
                class="text-indigo-600" href="/auth/signup">Sign Up&nbsp;<SmallArrowRight/></a></span>
    </aside>

    <main class="fixed top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2">
        <Card>
            <CardHeader slot="title">
                Sign In
            </CardHeader>
            <div class="flex flex-col gap-5 m-6 " slot="content">
                <Input bind:value={userIdentifier} error={!!userIdentifierError} label="Username/E-Mail"
                       on:input={clearError}
                       placeholder="Enter Username or E-Mail"
                       size="lg"/>
                <Input bind:value={password} error={!!passwordError} label="Password" on:input={clearError}
                       placeholder="Enter Password" size="lg"
                       type={showPassword ? "text": "password"}/>

                <Checkbox className="m-auto" onclick={handleShowHidePassword}>Show password</Checkbox>

                <Button className="m-auto" onClick={handleLogin}>Login</Button>

                <ErrorTextBox className="m-auto" bind:value={responseError}/>
            </div>

        </Card>
    </main>
</div>
