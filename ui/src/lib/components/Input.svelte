<script lang="ts">
    import { onMount } from "svelte";

    export let
        type: "button" | "checkbox" | "color" | "date" | "datetime-local" | "email" | "file" | "hidden" | "image" | "month" | "number" | "password" | "radio" | "range" | "reset" | "search" | "submit" | "tel" | "text" | "time" | "url" | "week"  = "text",
        placeholder: string,
        value: string,
        className: string,
        size = "sm"
    export let helperText: string;

    const sizes = new Map<string, string>([
        ["sm", "w-56"],
        ["md", "w-72"],
        ["lg", "w-80"],
        ["xl", "w-96"]
    ]);

    onMount(() => console.log(helperText));

    const handleInput = e => {
        value = e.target.value
    };
</script>

<div class="{className}">
    <p>{@html helperText}</p>
    <div class="flex justify-end">
        <span class="p-4 absolute">
            <slot />
        </span>
        <input type={type} value={value} placeholder={placeholder} on:input={handleInput} />
    </div>
    <p class="ml-1.5 mb-1 text-red font-bold text-left mt-2"></p>
</div>
