import "../styles/card.css";

import cardBackSvg from "../assets/Card_back_01.svg";

import CA from "../assets/C/A.svg";
import C2 from "../assets/C/2.svg";
import C3 from "../assets/C/3.svg";
import C4 from "../assets/C/4.svg";
import C5 from "../assets/C/5.svg";
import C6 from "../assets/C/6.svg";
import C7 from "../assets/C/7.svg";
import C8 from "../assets/C/8.svg";
import C9 from "../assets/C/9.svg";
import C10 from "../assets/C/10.svg";
import CJ from "../assets/C/J.svg";
import CQ from "../assets/C/Q.svg";
import CK from "../assets/C/K.svg";

import DA from "../assets/D/A.svg";
import D2 from "../assets/D/2.svg";
import D3 from "../assets/D/3.svg";
import D4 from "../assets/D/4.svg";
import D5 from "../assets/D/5.svg";
import D6 from "../assets/D/6.svg";
import D7 from "../assets/D/7.svg";
import D8 from "../assets/D/8.svg";
import D9 from "../assets/D/9.svg";
import D10 from "../assets/D/10.svg";
import DJ from "../assets/D/J.svg";
import DQ from "../assets/D/Q.svg";
import DK from "../assets/D/K.svg";

import HA from "../assets/H/A.svg";
import H2 from "../assets/H/2.svg";
import H3 from "../assets/H/3.svg";
import H4 from "../assets/H/4.svg";
import H5 from "../assets/H/5.svg";
import H6 from "../assets/H/6.svg";
import H7 from "../assets/H/7.svg";
import H8 from "../assets/H/8.svg";
import H9 from "../assets/H/9.svg";
import H10 from "../assets/H/10.svg";
import HJ from "../assets/H/J.svg";
import HQ from "../assets/H/Q.svg";
import HK from "../assets/H/K.svg";

import SA from "../assets/S/A.svg";
import S2 from "../assets/S/2.svg";
import S3 from "../assets/S/3.svg";
import S4 from "../assets/S/4.svg";
import S5 from "../assets/S/5.svg";
import S6 from "../assets/S/6.svg";
import S7 from "../assets/S/7.svg";
import S8 from "../assets/S/8.svg";
import S9 from "../assets/S/9.svg";
import S10 from "../assets/S/10.svg";
import SJ from "../assets/S/J.svg";
import SQ from "../assets/S/Q.svg";
import SK from "../assets/S/K.svg";

const ca = "ca";
const c2 = "c2";
const c3 = "c3";
const c4 = "c4";
const c5 = "c5";
const c6 = "c6";
const c7 = "c7";
const c8 = "c8";
const c9 = "c9";
const c10 = "c10";
const cj = "cj";
const cq = "cq";
const ck = "ck";

const da = "da";
const d2 = "d2";
const d3 = "d3";
const d4 = "d4";
const d5 = "d5";
const d6 = "d6";
const d7 = "d7";
const d8 = "d8";
const d9 = "d9";
const d10 = "d10";
const dj = "dj";
const dq = "dq";
const dk = "dk";

const ha = "ha";
const h2 = "h2";
const h3 = "h3";
const h4 = "h4";
const h5 = "h5";
const h6 = "h6";
const h7 = "h7";
const h8 = "h8";
const h9 = "h9";
const h10 = "h10";
const hj = "hj";
const hq = "hq";
const hk = "hk";

const sa = "sa";
const s2 = "s2";
const s3 = "s3";
const s4 = "s4";
const s5 = "s5";
const s6 = "s6";
const s7 = "s7";
const s8 = "s8";
const s9 = "s9";
const s10 = "s10";
const sj = "sj";
const sq = "sq";
const sk = "sk";

function getCard(id) {
  switch (id) {
    case sa:
      return SA;
    case s2:
      return S2;
    case s3:
      return S3;
    case s4:
      return S4;
    case s5:
      return S5;
    case s6:
      return S6;
    case s7:
      return S7;
    case s8:
      return S8;
    case s9:
      return S9;
    case s10:
      return S10;
    case sj:
      return SJ;
    case sq:
      return SQ;
    case sk:
      return SK;

    case ca:
      return CA;
    case c2:
      return C2;
    case c3:
      return C3;
    case c4:
      return C4;
    case c5:
      return C5;
    case c6:
      return C6;
    case c7:
      return C7;
    case c8:
      return C8;
    case c9:
      return C9;
    case c10:
      return C10;
    case cj:
      return CJ;
    case cq:
      return CQ;
    case ck:
      return CK;

    case da:
      return DA;
    case d2:
      return D2;
    case d3:
      return D3;
    case d4:
      return D4;
    case d5:
      return D5;
    case d6:
      return D6;
    case d7:
      return D7;
    case d8:
      return D8;
    case d9:
      return D9;
    case d10:
      return D10;
    case dj:
      return DJ;
    case dq:
      return DQ;
    case dk:
      return DK;

    case ha:
      return HA;
    case h2:
      return H2;
    case h3:
      return H3;
    case h4:
      return H4;
    case h5:
      return H5;
    case h6:
      return H6;
    case h7:
      return H7;
    case h8:
      return H8;
    case h9:
      return H9;
    case h10:
      return H10;
    case hj:
      return HJ;
    case hq:
      return HQ;
    case hk:
      return HK;

    default:
      return <div>not found</div>;
  }
}

function Card({ id }) {
  return (
    <div className="scene scene--card">
      <div className="card">
        <div className=" card__face cardBack">
          <img
            src={getCard(id)}
            style={{ boxShadow: " -2px 1px 4px -1px black" }}
            className="cardSize"
            alt="card front "
          />
        </div>
        <div className="card__face  cardFront ">
          <img
            src={cardBackSvg}
            style={{ boxShadow: " -2px 1px 4px -1px black" }}
            className="cardSize"
            alt="card back "
          />
        </div>
      </div>
    </div>
  );
}

export default Card;
