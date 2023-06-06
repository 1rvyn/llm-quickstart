import os
from langchain.vectorstores.chroma import Chroma
from langchain.chat_models import ChatOpenAI
from langchain.embeddings.openai import OpenAIEmbeddings
from langchain.chains import ConversationalRetrievalChain
from langchain.schema import HumanMessage, AIMessage
from dotenv import load_dotenv
from flask import Flask, request, jsonify

app = Flask(__name__)

load_dotenv()

def make_chain():
    model = ChatOpenAI(
        model_name="gpt-3.5-turbo", 
        temperature="0",
        # verbose=True
    )
    embedding = OpenAIEmbeddings()

    vector_store = Chroma(
        collection_name="june-2023-quickstartsimulator",
        embedding_function=embedding,
        persist_directory="src/data/chroma",
    )

    return ConversationalRetrievalChain.from_llm(
        model,
        retriever=vector_store.as_retriever(),
        return_source_documents=True,
        # verbose=True,
    )

chain = make_chain()
chat_history = []

@app.route('/ask', methods=['POST'])
def ask():

    data = request.get_json()
    question = data['question']
    # question = request.form['question']

    # Generate answer
    response = chain({"question": question, "chat_history": chat_history})

    # Retrieve answer
    answer = response["answer"]
    source = response["source_documents"]
    if source:
        page_numbers = set(doc.metadata['page_number'] for doc in source)
        page_numbers_str = ', '.join(str(pn) for pn in page_numbers)
        answer += f"\n\nYou can read about this on page {page_numbers_str} on our quick-start guide."
    
    chat_history.append(HumanMessage(content=question))
    chat_history.append(AIMessage(content=answer))

    # Return answer
    return jsonify({'answer': answer})

if __name__ == "__main__":
    load_dotenv()

    app.run(threaded=True, port=5000)

#   if source:
#         source_str = ', '.join(str(doc) for doc in source)
#         answer += f"\n\nYou can read about this on page {source_str} on our quick-start guide."
    