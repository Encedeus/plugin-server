<script lang="ts">

    export let
        type: "button" | "checkbox" | "color" | "date" | "datetime-local" | "email" | "file" | "hidden" | "image" | "month" | "number" | "password" | "radio" | "range" | "reset" | "search" | "submit" | "tel" | "text" | "time" | "url" | "week" = "text",
        placeholder: string,
        value: string,
        className: string,
        size = "sm";
    export let helperText: string;
    export let inline: boolean;

    const handleInput = e => {
        value = e.target.value;
    };
</script>

<div class="{className}">
    <p class:inlineComponent={inline}>{@html helperText}</p>
    <div class:inlineComponent={inline}>
        <span class:inlineComponent={inline}>
            <slot/>
        </span>
        <input class:inlineComponent={inline}
               type={type}
               value={value}
               placeholder={placeholder}
               on:input={handleInput}
        />
    </div>
    <p></p>
</div>

<style>
    .inlineComponent {
        display: inline;
    }
</style>