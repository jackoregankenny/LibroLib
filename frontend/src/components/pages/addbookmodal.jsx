import React, { useState } from 'react';
import {
    Modal,
    ModalOverlay,
    ModalContent,
    ModalHeader,
    ModalFooter,
    ModalBody,
    ModalCloseButton,
    Button,
    FormControl,
    FormLabel,
    Input
} from '@chakra-ui/react';
import { runtime } from '@wails/runtime';

const AddBookModal = ({ isOpen, onClose }) => {
    const [bookData, setBookData] = useState({
        Title: '',
        Author: '',
        Genre: '',
        PublicationDate: '',
        Publisher: '',
        Language: '',
        ISBN: '',
        PageCount: '',
        Read: false,
        Rating: '',
        Notes: '',
        CoverImagePath: '',
        FilePath: '', // Path of the uploaded file
    });

    // Handle manual input change
    const handleInputChange = (e) => {
        setBookData({ ...bookData, [e.target.name]: e.target.value });
    };

    const handleFileChange = async (event) => {
        const file = event.target.files[0];
        if (file) {
            const filePath = await uploadFileAndGetPath(file);

            if (file.name.endsWith('.epub')) {
                runtime.Invoke("ExtractEPUBMetadata", filePath)
                    .then((metadata) => {
                        setBookData({ ...bookData, ...metadata, FilePath: filePath });
                    })
                    .catch(console.error);
            } else {
                setBookData({ ...bookData, FilePath: filePath });
            }
        }
    };

    // Handle form submission
    const handleSubmit = () => {
        runtime.Invoke("AddBookToLibrary", bookData)
            .then(() => {
                // Handle success - close the modal
                onClose();
            })
            .catch(console.error);
    };

    return (
        <>
            <Button onClick={onOpen}>Add Book</Button>
            <Modal isOpen={isOpen} onClose={onClose}>
                <ModalOverlay />
                <ModalContent>
                    <ModalHeader>Add a New Book</ModalHeader>
                    <ModalCloseButton />
                    <ModalBody pb={6}>
                        {/* Input fields for each metadata attribute */}
                        <FormControl>
                            <FormLabel>Title</FormLabel>
                            <Input name="Title" value={bookData.Title} onChange={handleInputChange} />
                        </FormControl>
                        <FormControl mt={4}>
                            <FormLabel>Author</FormLabel>
                            <Input name="Author" value={bookData.Author} onChange={handleInputChange} />
                        </FormControl>
                        <FormControl>
                            <FormLabel>Genre</FormLabel>
                            <Input name="Genre" value={bookData.Genre} onChange={handleInputChange} />
                        </FormControl>
                        <FormControl>
                            <FormLabel>Publication Date</FormLabel>
                            <Input name="PublicationDate" value={bookData.PublicationDate} onChange={handleInputChange} />
                        </FormControl>
                        {/* Repeat for other fields like PublicationDate, Publisher, Language, ISBN, PageCount */}
                        <FormControl mt={4}>
                            <FormLabel>Upload EPUB File</FormLabel>
                            <Input type="file" accept=".epub" onChange={handleFileChange} />
                        </FormControl>
                    </ModalBody>
                    <ModalFooter>
                        <Button colorScheme="blue" mr={3} onClick={handleSubmit}>
                            Add Book
                        </Button>
                        <Button onClick={onClose}>Cancel</Button>
                    </ModalFooter>
                </ModalContent>
            </Modal>
        </>
    );
};

export default AddBookModal;
