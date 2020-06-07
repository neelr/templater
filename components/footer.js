import React from "react"
import { Flex, Text, Link } from "rebass"

const A = props => <Link sx={{
    color: "primary",
    ":hover": {
        color: "secondary",
        textDecorationStyle: "wavy"
    }
}} {...props} />

export default () => (
    <Flex mt="auto" height="15vh">
        <Text m="auto">Open Source | MIT | <A href="https://github.com/neelr/templater">CLI</A> | <A href="https://github.com/neelr/templater/tree/website">Source Code</A></Text>
    </Flex>
)