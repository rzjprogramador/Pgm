const newUserRequest1 = { nome: "Reinaldo", sobrenome: "Zacharias"}

const createUser = (newUser) => newUser

const UserRepositoryMemory = {
  create (newUser) {
    return newUser
  }
}

console.log(createUser(newUser1))
