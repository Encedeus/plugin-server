<script lang="ts">
    import {goto} from "$app/navigation";

    export let size: "sm" | "md" | "lg" | "xl" = "md";
    export let color: "indigo" | "red" = "indigo";
    export let type: "button" | "submit" = "button";
    export let className = "";
    export let onClick: () => void = () => {};
    export let redirect: string;
    export let isDisabled: boolean = false;

    const sizes = new Map<string, string>([
        ["xs", "w-20 h-5.5 font-normal"],
        ["sm", "w-24 h-8"],
        ["md", "w-32 h-11"],
        ["lg", ""],
        ["xl", ""],
    ]);

    const colors = new Map<string, string>([
        ["indigo", "bg-indigo-600 hover:bg-indigo-700 active:bg-indigo-800 disabled:bg-indigo-500"],
        ["red", "bg-red-600 hover:bg-red-700 active:bg-red-800"],
    ]);

    async function handleClick() {
        await goto(redirect);
        onClick()
    }

    let disabled = isDisabled ? "disabled" : "";
    console.log(isDisabled);
</script>

{#if redirect}
    <button on:click={handleClick}
            {type}
            {disabled}
            class="text-white font-bold text-sm rounded-full {colors.get(color)} {sizes.get(size)} hover:shadow-xl active:shadow-xl transition-all {className}"
    >
        <slot/>
    </button>
{:else }
    <button on:click={onClick}
            {type}
            {disabled}
            class="text-white font-bold text-sm rounded-full {colors.get(color)} {sizes.get(size)} hover:shadow-xl active:shadow-xl transition-all {className}"
    >
        <slot/>
    </button>
{/if}