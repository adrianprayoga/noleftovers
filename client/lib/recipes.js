export const getAllRecipes = () => {
    return sampleDataWithId
}

const sampleData = [
    {
        name: 'smash burger',
        description: `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`
    },
    {
        name: 'oklahoma burger',
        description: `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged.`
    },
    {
        name: 'neapolitan pizza',
        description: `Lorem Ipsum is simply dummy text of the printing and typesetting industry. `
    },
    {
        name: 'neapolitan pizza',
        description: `Lorem Ipsum is simply dummy text of the printing and typesetting industry. `
    },
    {
        name: 'neapolitan pizza',
        description: `Lorem Ipsum is simply dummy text of the printing and typesetting industry. `
    },
    {
        name: 'neapolitan pizza',
        description: `Lorem Ipsum is simply dummy text of the printing and typesetting industry. `
    },
    {
        name: 'neapolitan pizza',
        description: `Lorem Ipsum is simply dummy text of the printing and typesetting industry. `
    }
]

const sampleDataWithId = sampleData.map((recipe, i) => {
    return {
        ...recipe,
        id: String(i)
    }
})


export function getAllRecipeIds() {
    return sampleDataWithId.map(recipe => {
        return {
            params: {
                id: recipe.id,
            },
        };
    });
}


export const getRecipeData = async (id) => {
    const data = sampleDataWithId.find(recipe => recipe.id === id);

    return {
        ...data
    }
}