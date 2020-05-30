from sqlalchemy.dialects.mysql import VARCHAR, INTEGER
from sqlalchemy import Column
from variables import DbVars
from sqlalchemy.orm import relationship, joinedload
from sqlalchemy import ForeignKey
from base import BASE
from db import transaction_context, create_table
from pprint import pprint as pp


class Parent(BASE):

    __tablename__ = 'parent'

    __table_args__ = {'mysql_engine': DbVars.ENGINE,
                      'mysql_charset': DbVars.CHARSET,
                      'mysql_collate': DbVars.COLLATE}

    id = Column(INTEGER(unsigned=True), primary_key=True, autoincrement=True)

    name = Column(VARCHAR(128), unique=True)

    childs = relationship('Child',
                          cascade='all, delete-orphan',
                          passive_deletes=True,
                          back_populates='parent',
                          lazy='select')


class Child(BASE):

    __tablename__ = 'child'

    __table_args__ = {'mysql_engine': DbVars.ENGINE,
                      'mysql_charset': DbVars.CHARSET,
                      'mysql_collate': DbVars.COLLATE}

    id = Column(INTEGER(unsigned=True), primary_key=True, autoincrement=True)

    parent_id = Column(INTEGER(unsigned=True),
                       ForeignKey('parent.id', ondelete="CASCADE"))

    name = Column(VARCHAR(256), unique=True)

    parent = relationship('Parent',
                          back_populates='childs')


create_table()

def insert_test_data():
    with transaction_context() as session:
        p1 = Parent()
        p1.name = "p1"

        c1 = Child()
        c1.name = "c1"
        c2 = Child()
        c2.name = "c2"

        p1.childs = [c1, c2]

        session.add(p1)


def test_delete_cascade():
    with transaction_context() as session:
        session.query(Parent).filter(Parent.name == "p1").delete()


def test_select():
    with transaction_context() as session:
        p1 = session.query(Parent).options(joinedload("childs")).filter(Parent.name == "p1").one_or_none()
        print(p1.name)
        for child in p1.childs:
            print(child.name)


def test():
    insert_test_data()
    test_delete_cascade()
    test_select()


if __name__ == "__main__":
    test()