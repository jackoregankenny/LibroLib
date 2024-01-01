import React from 'react';
import { Flex, IconButton, useDisclosure, Box, Link } from '@chakra-ui/react';
import { HamburgerIcon } from '@chakra-ui/icons';
import { Link as RouterLink } from 'react-router-dom';

function Header() {
  const { isOpen, onOpen, onClose } = useDisclosure();

  const SidebarContent = () => (
    <Box>
      <Link as={RouterLink} to="/" p="4">Home</Link>
      <Link as={RouterLink} to="/library" p="4">Library</Link>
      {/* Additional Links */}
    </Box>
  );

  return (
    <Flex as="header" align="center" justify="space-between" p="4" bg="black" color="white">
      <IconButton
        icon={<HamburgerIcon />}
        aria-label="Open Menu"
        onClick={onOpen}
        colorScheme="whiteAlpha"
      />
      {/* Drawer Component for Mobile Navigation */}
      {/* Desktop Navigation */}
      <Flex display={{ base: 'none', md: 'flex' }}>
        <SidebarContent />
      </Flex>
    </Flex>
  );
}

export default Header;
