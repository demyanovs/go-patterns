const allChannels = '/channels/all'
const allowedChannels = '/channels/allowed'

const getAllChannels = async () => {
    fetch(allChannels)
}

const getAllowedChannels = async () => {
    fetch(allowedChannels)
}

const processChannels = async () => {
    const allowedChannels = await getAllowedChannels()
    const allChannels = await getAllChannels()

    // 'animalPlanetCartoonsMovies'

    // ['animalPlanet', 'Movies']
}